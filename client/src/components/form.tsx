import React, { FC } from "react";
import { useFormik } from "formik";
import Majors from "../data/majors";

export interface FormValues {
  major: string;
}

interface Props {
  handleSubmit: (values: FormValues) => void;
}

export const CourseForm: FC<Props> = ({ handleSubmit }) => {
  const formik = useFormik({
    initialValues: { major: "DS25" },
    onSubmit: handleSubmit,
  });

  return (
    <div>
      <form
        onSubmit={formik.handleSubmit}
        style={{ display: "flex", flexDirection: "column" }}
      >
        <label htmlFor="major" style={{ fontFamily: 'SF Pro Display', fontWeight: 250, }}>Major</label>
        <select name="major" onChange={formik.handleChange}>
          {Object.entries(Majors).map(([name, code]) => (
            <option value={code}>{name}</option>
          ))}
        </select>

        <button type="submit">Submit</button>
      </form>
    </div>
  );
};
