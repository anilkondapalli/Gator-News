import React from "react";
import User from './User';
import { BrowserRouter as Router } from 'react-router-dom'; 
import { render } from '@testing-library/react';


describe('User', () => {
    test('Render User Page', () => {
        <Router>
            render(<User/>);
            </Router>
    });

    test("test_Basic_Layout_ClassName",()=>{
        const {container} = render(
            <Router>
                render(<User/>);
            </Router>
        );
        // eslint-disable-next-line testing-library/no-node-access
        expect(container.firstChild).not.toEqual("user");

       });
});