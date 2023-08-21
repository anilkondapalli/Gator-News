import { within, render, screen } from '@testing-library/react';
import React from "react";
import Users from './Users';
import AllUsers from './Users';
import { Router } from 'react-router-dom';

import { MemoryRouter } from "react-router-dom";


describe('AllUsers', () => {

   
    it("should render title", () => {
        const view = render(
            <MemoryRouter>
              <AllUsers/>
           </MemoryRouter>
          );
        expect(screen.queryAllByText("Users")).toHaveLength(2);
        
      });
      
      test("Search Users field should have Placeholder",()=>{
        const view = render(
            <MemoryRouter>
              <AllUsers/>
           </MemoryRouter>
          );        
          screen.queryByPlaceholderText(/Search Users/i)
      });

      test("test_Users Initial Component",()=>{
        const {container} = render(
            <MemoryRouter>
              <AllUsers/>
            </MemoryRouter>
        );
        // eslint-disable-next-line testing-library/no-node-access
        expect(container.firstChild).toHaveClass('row');

       });

       test("test_Users Ending Component",()=>{
        const {container} = render(
            <MemoryRouter>
              <AllUsers/>
            </MemoryRouter>
        );
        // eslint-disable-next-line testing-library/no-node-access
        expect(container.lastChild).toHaveClass('row');

       });
});