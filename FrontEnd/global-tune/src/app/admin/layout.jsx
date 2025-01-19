import React from "react";
import AdminNavbar from '@/components/AdminNavBar'

function layout({ children }) {
  return (
    <>
      <header>
    <AdminNavbar/>
      </header>
      {children}
    </>
  );
}

export default layout;
