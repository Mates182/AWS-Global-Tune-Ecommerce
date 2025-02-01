"use client";
import React from "react";
import Link from "next/link";
import { useEffect } from "react";

function AdminNavbar() {
  useEffect(() => {
    require("bootstrap/dist/js/bootstrap.bundle.min.js");
  }, []);
  return (
    <div className="bg-light">
      <nav className="container navbar navbar-expand-lg navbar-light">
        <Link href="/admin" className="navbar-brand">
          GlobalTune Admin
        </Link>
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarTogglerDemo02"
          aria-controls="navbarTogglerDemo02"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>

        <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
          <ul className="navbar-nav me-auto mb-2 mb-lg-0">
            <li className="nav-item active">
              <Link href="/admin" className="nav-link">
                Home
              </Link>
            </li>
            <li className="nav-item">
              <Link
                href="#"
                className="nav-link disabled"
                tabIndex="-1"
                aria-disabled="true"
              >
                My Account
              </Link>
            </li>
          </ul>
          <form className="d-flex">
            <input
              className="form-control me-2"
              type="search"
              placeholder="Search"
              aria-label="Search"
            />
            <button
              className="btn btn-outline-success"
              type="submit"
              onClick={(e) => {
                e.preventDefault();
                alert("searching");
              }}
            >
              Search
            </button>
            <button
              className="btn btn-outline-danger"
              onClick={(e) => {
                e.preventDefault();
                alert("loggin out");
              }}
            >
              LogOut
            </button>
          </form>
        </div>
      </nav>

      <hr className="container mt-0 mb-0" />

      <nav className="container navbar navbar-expand-lg navbar-light">
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarDropdownMenu"
          aria-controls="navbarDropdownMenu"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarDropdownMenu">
          <ul className="navbar-nav me-auto mb-2 mb-lg-0">
            <li className="nav-item dropdown">
              <Link
                href="#"
                className="nav-link dropdown-toggle"
                id="navbarDropdownMenuLink"
                role="button"
                aria-expanded="false"
              >
                Products
              </Link>
              <ul
                className="dropdown-menu"
                aria-labelledby="navbarDropdownMenuLink"
              >
                <li>
                  <Link href="/admin/dashboard" className="dropdown-item">
                    Dashboard
                  </Link>
                </li>
              </ul>
            </li>
          </ul>
        </div>
      </nav>
    </div>
  );
}

export default AdminNavbar;
