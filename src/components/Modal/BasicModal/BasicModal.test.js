import { within, render, screen } from '@testing-library/react';
import React from "react";
import BasicModal from './BasicModal';
import Gator from '../../../assets/png/logo-white.png';
import { Modal } from "react-bootstrap";


describe('BasicModal', () => {

   
    test('Modal renders Modal header title body', () => {
        render(<BasicModal />);
        const appHeader = screen.queryByTestId('modal-header');
   
        console.log(appHeader.className);
      });

      test('DOM', () => {
        new Promise(
          (resolve, reject) => {
            let componentResult = screen.queryAllByTestId('modal-header')
            componentResult instanceof Error? reject(componentResult):resolve(componentResult)
          })
          .then(() => {
            console.log('Component Found');
          })
          .catch((e) => {
            console.log('Error in Capturing: ' + e.message);
          });
      });

      

});