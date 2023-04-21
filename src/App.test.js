import { fireEvent, render, screen } from "@testing-library/react";
import App from "./App";

test("renders app", () => {
  render(<App />);
  const linkElement = screen.getByText(/bridge 1/i);
  expect(linkElement).toBeInTheDocument();
});

test("click add bridge", () => {
  render(<App />);
  fireEvent.click(screen.getByTestId(/add-bridge/));

  let linkElement = screen.getByText(/bridge 2/i);
  expect(linkElement).toBeInTheDocument();

  fireEvent.click(screen.getByTestId(/add-bridge/));

  linkElement = screen.getByText(/bridge 3/i);
  expect(linkElement).toBeInTheDocument();
});

test("click add hiker", () => {
  render(<App />);
  fireEvent.click(screen.getByTestId(/add-hiker/));

  let linkElement = screen.getByText(/hiker 1/i);
  expect(linkElement).toBeInTheDocument();

  fireEvent.click(screen.getByTestId(/add-hiker/));

  linkElement = screen.getByText(/hiker 2/i);
  expect(linkElement).toBeInTheDocument();
});
