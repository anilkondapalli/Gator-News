import { fireEvent, render, screen } from '@testing-library/react';
import React from "react";
import { MemoryRouter } from 'react-router-dom/cjs/react-router-dom.min';
import ListTweets from './ListTweets';


test("test_tweeetsClassName",()=>{
 const {container} = render(
     <MemoryRouter>
        <ListTweets />
     </MemoryRouter>
 );
 // eslint-disable-next-line testing-library/no-node-access
 expect(container.firstChild).toHaveClass('list-tweets');
});