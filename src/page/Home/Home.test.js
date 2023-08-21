import React from "react";
import Home from './Home';
import { BrowserRouter as Router } from 'react-router-dom'; 

import {render } from '@testing-library/react';
import { Link } from 'react-router-dom';
import { StaticRouter } from 'react-router'
import renderer from 'react-test-renderer';

describe('Home', () => {
    test('Render Home Page', () => {
        <Router>
            render(<Home/>);
           </Router>
    });

    
});