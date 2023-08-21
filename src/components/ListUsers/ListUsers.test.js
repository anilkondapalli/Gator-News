import React from "react";
import ListUsers from './ListUsers';
import { BrowserRouter as Router } from 'react-router-dom'; 
import { render } from '@testing-library/react';


describe('ListUsers', () => {
    test('Render ListUsers Page', () => {
       
            render(<ListUsers/>);
           
    });

    test("test List Users Components",()=>{
        const {container} = render(
            <ListUsers/>
        );
        // eslint-disable-next-line testing-library/no-node-access
        expect(container.firstChild).toContainHTML('<h2>No Results found</h2>');

       });
});