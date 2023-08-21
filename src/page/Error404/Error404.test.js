
import React from "react";
import Error404 from './Error404';
import { BrowserRouter as Router } from 'react-router-dom'; 
import Error404Image from "../../assets/png/404-Error.png";
import { Link } from 'react-router-dom';
import { StaticRouter } from 'react-router'
import renderer from 'react-test-renderer';

// Since Link cant be used outside router we have to rap all the calls with <Router>
describe('Error404', () => {
    test('Render Error404 Page', () => {
        <Router>
            render(<Error404/>);
        </Router> 
    });

    test('Error404 must have src={Logo} and alt="GatorLogo"', () => {
        <Router>
            render(<Error404/>);
            const logo = screen.getByRole('img');
            expect(logo).toHaveAttribute('src', Logo);
            expect(logo).toHaveAttribute('alt', 'GatorLogo');
        </Router>
      });

      test('Error404 must have src={Error404Image} and alt="Error404', () => {
        <Router>
            render(<Error404/>);
            const logo = screen.getByRole('img');
            expect(logo).toHaveAttribute('src', Error404Image);
            expect(logo).toHaveAttribute('alt', 'Error404');
        </Router>
      });

      test('Link in Error 404 matches snapshot', () => {
        const error404Page = renderer.create(
          <StaticRouter>
            <Link to="/" />
          </StaticRouter>
        );
      
        let componentTree = error404Page.toJSON();
        expect(componentTree).toMatchSnapshot();
      });
});