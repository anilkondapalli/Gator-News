
import { fireEvent, render, screen } from '@testing-library/react';
import React from "react";
import { MemoryRouter } from 'react-router-dom/cjs/react-router-dom.min';
import LeftMenu from './LeftMenu';

test ("home",()=>{
    render(
    <MemoryRouter>
    <LeftMenu />
    </MemoryRouter>)
    const linkelement = screen.getByText('Home')
    expect(linkelement).toBeInTheDocument();
});

test ("users",()=>{
    render(
    <MemoryRouter>
    <LeftMenu />
    </MemoryRouter>)
    const linkelement = screen.getByText('Users')
    expect(linkelement).toBeInTheDocument();
});

test ("profile",()=>{
    render(
    <MemoryRouter>
    <LeftMenu />
    </MemoryRouter>)
    const linkelement = screen.getByText('Profile')
    expect(linkelement).toBeInTheDocument();
});

test ("logout",()=>{
    render(
    <MemoryRouter>
    <LeftMenu />
    </MemoryRouter>)
    const linkelement = screen.getByText('Logout')
    expect(linkelement).toBeInTheDocument();
});
