import React, { FC } from 'react';
import styled from '@emotion/styled';
import { Header, Subheader } from '../components/header'
import Plane from '../components/plane'
import { keyframes } from '@emotion/react';

const Container = styled.div({
    display: 'flex',
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
    flexDirection: 'column'
})

const flyAnimation = keyframes`
    0%, 100% { transform: translateY(5px); };
    50% { transform: translateY(-5px); };
`

const PlaneContainer = styled.div({
    zoom: 4,
    animation: `${flyAnimation} 5s ease infinite`,
    padding: 20,
})

const PlaneText = styled(Subheader)({
    fontStyle: 'italic'
})

const LoadingPage: FC = () => {
    return (
        <Container>
            <Header/>
            <Subheader>
                Welcome to your course destination! We’re glad to have you on board.
            </Subheader>

            <PlaneContainer>
                <Plane />
            </PlaneContainer>
            <PlaneText><i>Testing the winds...</i></PlaneText>
        </Container>
    )
}

export default LoadingPage;