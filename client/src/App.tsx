import React, { FC, useState } from "react";
import LandingPage from "./pages/landing_page";
import LoadingPage from "./pages/loading_page";
import CoursePage from "./pages/course_page";
import { FormValues } from "./components/form";
import { CourseSchedule } from "./types";

const App: FC = () => {
  const [curOption, setCurOption] = useState<FormValues>();
  const [schedule, setSchedule] = useState<CourseSchedule>();
  const [loading, setLoading] = useState(false);

  const handleFormSubmit = async (values: FormValues) => {
    setCurOption(values);
    setLoading(true);
    const resp = await fetch(
      `http://localhost:8010/course_plan/${values.major}`
    );
    const schedule: CourseSchedule = await resp.json();
    setSchedule(schedule);
    setLoading(false);
  };

  const refresh = async () => {
      if (!curOption) {
          throw new Error('Need current options to refresh')
      }

      handleFormSubmit(curOption);
  }

  if (loading) {
    return <LoadingPage />;
  }

  if (schedule) {
    return <CoursePage schedule={schedule} refresh={refresh} />;
  }

  return <LandingPage handleSubmit={handleFormSubmit} />;
};

export default App;
