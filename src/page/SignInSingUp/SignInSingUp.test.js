import React from "react";
import SignInSingUp from './SignInSingUp';
import { BrowserRouter as Router } from 'react-router-dom'; 
import { render } from '@testing-library/react';


describe('SignInSingUp', () => {
    test('Render SignInSingUp Page', () => {
        <Router>
            render(<SignInSingUp/>);
           </Router>
    });

    test("test_signin-signupClassName",()=>{
        const {container} = render(
               <SignInSingUp />
        );
        // eslint-disable-next-line testing-library/no-node-access
        expect(container.firstChild).toHaveClass('signInSignUp');

       });
});