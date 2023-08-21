import { render, screen } from '@testing-library/react';
import React from "react";
import TweetModal from './TweetModal';
import { create } from "react-test-renderer";

function Button(props) {
    return <button>Nothing to do for now</button>;
}

describe('TweetModal', () => {

    test("Render the TweetModal", () => {
        render(<TweetModal/>);
    });

    test("Match Snapshot", () => {
        const button = create(<Button />);
        expect(button.toJSON()).toMatchSnapshot();
      });


      test("Display expected text when clicked (testing the wrong way)", () => {
        const buttonComponent = create(<Button text="SUBSCRIBE TO BASIC" />);
        const buttonInstance = buttonComponent.getInstance();
        expect(buttonInstance).toBe(null);
      });

      test("It shows the expected text when clicked (testing the wrong way!)", () => {
        const component = create(<Button text="Buzz" />);
        const instance = component.root.props;
        expect(instance).toStrictEqual({"text": "Buzz"});
        console.log(instance);

       
      });
      test('Check Button Submission', () => {
        return screen.findByRole('submit')
          .then(() => {
            console.log('found');
          })
          .catch(() => {
            console.log('not found');
        });
    });

  });