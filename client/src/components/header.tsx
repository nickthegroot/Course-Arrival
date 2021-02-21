import React, { FC } from 'react'
import styled from '@emotion/styled'
import PlaneTakeoff from './plane_takeoff'

export const Heading = styled.h1({
    fontFamily: 'SF Pro Display',
    fontSize: 40,
    textAlign: 'center',
    textColor: 'black',
    fontWeight: 'bold',
})

export const Subheading = styled.h2({
    fontFamily: 'SF Pro Display',
    fontSize: 20,
    textAlign: 'center',
    textColor: 'black',
    fontWeight: 250,
})

const HeaderContainer = styled.div({
    display: 'flex',
    width: '100%',
    alignItems: 'center',
    justifyContent: 'center',
})

export const Header: FC = () => (
    <HeaderContainer>
        <PlaneTakeoff style={{ padding: 10 }} />
        <Heading>Course Arrival</Heading>
    </HeaderContainer>
)

export const Subheader: FC = ({ children }) => (
    <HeaderContainer>
        <Subheading>{children}</Subheading>
    </HeaderContainer>
)