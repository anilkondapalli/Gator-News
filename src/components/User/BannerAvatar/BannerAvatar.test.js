import { fireEvent, render, screen } from '@testing-library/react';
import React from "react";
import Routing from './BannerAvatar';
import banner from '../../../assets/png/logo.png';
import avatar from "../../../assets/png/user.png";


describe('Routing', () => {
    test('BannerAvatar must have src = {banner} and alt = "Gator"', () => {
      render(<Routing/>);
      const logo =screen.getByTestId("banner_avatar");
      expect(logo).toHaveAttribute('src', banner);
      expect(logo).toHaveAttribute('alt', 'Gator');
    });

    test('UserAvatar must have src={avatar} and alt="UserAvatar"', () => {
        render(<Routing/>);
        const logo =screen.getByTestId("user_avatar");
        expect(logo).toHaveAttribute('src', avatar);
        expect(logo).toHaveAttribute('alt', 'UserAvatar');
      });

    test("Should render the Routing", () => {
        render(<Routing/>);
    });
 

  });