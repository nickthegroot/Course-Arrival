import React, { FC } from 'react';
import styled from '@emotion/styled';
import { Header, Subheader } from '../components/header'
import { CourseForm, FormValues } from '../components/form'

const Container = styled.div({
    display: 'flex',
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    flexDirection: 'column'
})

interface Props {
    handleSubmit: (values: FormValues) => void;
  }

const LandingPage: FC<Props> = ({ handleSubmit }) => {
    return (
        <Container>
            <Header/>
            <Subheader>
                Welcome to your course destination! We’re glad to have you on board.<br />
                There’s just a little bit of information we need before take-off.
            </Subheader>

            <CourseForm handleSubmit={handleSubmit} />
        </Container>
    )
}

export default LandingPage;