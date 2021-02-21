
export interface Course {
    id: string;
    title: string;
    units: string;
    description: string;
    upperDivOnly: boolean;
    gradOnly: boolean;
}

export interface CourseSchedule {
    course_plan: Course[];
    dot: string;
}