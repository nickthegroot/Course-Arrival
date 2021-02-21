import React, { FC } from 'react';
import { CourseSchedule } from '../types';

interface Props {
    schedule: CourseSchedule
}

const CoursePlan: FC<Props> = ({ schedule }) => (
    <>
    {schedule.course_plan.map(course => (
        <div style={{
            width: '80%',
            backgroundColor: 'white',
            boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19)',
            margin: 10,
            padding: 10,
            fontFamily: 'SF Pro Display',
        }}>
            <p>{course.id}: <span style={{ fontWeight: 100 }}>{course.title}</span></p>
            <p style={{ fontWeight: 100 }}>{course.description}</p>
        </div>
    ))}
    </>
)

export default CoursePlan;