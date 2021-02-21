import styled from '@emotion/styled';

export const Button = styled.button({
    backgroundColor: '#48E5C2',
    boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19)',
    width: '15vw',
    margin: 10,
    padding: 10,
    fontFamily: 'SF Pro Display',
    borderWidth: 0,
    transition: '0.5s',
    ':hover': {
        boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)',
    }
    // margin: 5,
})