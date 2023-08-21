import { fireEvent, render, screen } from '@testing-library/react';
import React from "react";
import InfoUser from './InfoUser';

test("Should be able to submit the User Info", () => {
    const mockFunction = jest.fn();
    render(<InfoUser handleSubmit={mockFunction} />);
  });

  test("Should render the User Info", () => {
    render(<InfoUser/>);
    });

