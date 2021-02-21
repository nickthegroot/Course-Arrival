import React, { FC } from 'react';
import styled from '@emotion/styled';
import { Graphviz } from 'graphviz-react';
import { Header, Subheader } from '../components/header'
import { CourseSchedule } from '../types';
import CoursePlan from '../components/course_plan';
import { Button } from '../components/button';

const Container = styled.div({
    display: 'flex',
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    flexDirection: 'column'
})

const DotContainer = styled.div({
    zoom: 2,
    boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19)',
    '.graph > polygon' : {
        fill: '#FCFAF9'
    },
})

interface Props {
    schedule: CourseSchedule;
    refresh: () => void;
}

const CoursePage: FC<Props> = ({ schedule, refresh }) => {
    return (
        <Container>
            <Header/>
            <Subheader>
                Welcome to your course destination! We’re glad to have you on board.
            </Subheader>

            
            <CoursePlan schedule={schedule} />
            <Button onClick={refresh}>Refresh</Button>

            <br />
            <DotContainer>
                <Graphviz dot={schedule.dot} />
            </DotContainer>
        </Container>
    )
}

export default CoursePage;