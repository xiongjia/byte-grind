import React from "react";

export const Button: React.FC<React.PropsWithChildren> = ({ children }) => {
  return (
    <button style={{ padding: "8px 16px", background: "#646cff", color: "#fff", border: "none", borderRadius: "6px" }}>
      {children}
    </button>
  );
};
