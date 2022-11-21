import React, { MouseEventHandler } from "react";

type ButtonType = {
  children: React.ReactNode;
  color?: string;
  onClick?: MouseEventHandler;
  type?: "button" | "submit";
};

function BlueButton({ children, ...props }: ButtonType) {
  return (
    <button
      type="button"
      className={`inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out`}
      {...props}
    >
      {children}
    </button>
  );
}

function RedButton({ children, ...props }: ButtonType) {
  return (
    <button
      type="button"
      className={`inline-block px-6 py-2.5 bg-red-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-red-700 hover:shadow-lg focus:bg-red-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-red-800 active:shadow-lg transition duration-150 ease-in-out`}
      {...props}
    >
      {children}
    </button>
  );
}

export default function Button({ children, color, ...props }: ButtonType) {
  if (color === "red") {
    return <RedButton {...props}>{children}</RedButton>;
  }
  return <BlueButton {...props}>{children}</BlueButton>;
}
