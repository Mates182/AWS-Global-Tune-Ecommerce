import React from "react";
// TODO: use our data on the comboxes
// add payment methods
// Redirect to purchase complete page
function BillingDetails() {
  return (
    <div>
      <form>
        <h2>Billing Details</h2>
        <div className="mb-3">
          <label className="form-label" htmlFor="personType">
            Person Type <span>*</span>
          </label>
          <select
            id="personType"
            className="form-select"
            name="personType"
            required
          >
            <option value="Natural">Individual</option>
            <option value="Legal">Company</option>
          </select>
        </div>

        <div className="mb-3">
          <label className="form-label" htmlFor="idType">
            Identification Type <span>*</span>
          </label>
          <select className="form-select" id="idType" name="idType" required>
            <option value="ID">ID</option>
            <option value="Passport">Passport</option>
          </select>
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="idNumber">
            Identification Number <span>*</span>
          </label>
          <input
            type="text"
            id="idNumber"
            className="form-control"
            name="idNumber"
            placeholder="Enter your ID number"
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="firstName">
            First Name <span>*</span>
          </label>
          <input
            className="form-control"
            type="text"
            id="firstName"
            name="firstName"
            placeholder="Enter your first name"
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="lastName">
            Last Name <span>*</span>
          </label>
          <input
            className="form-control"
            type="text"
            id="lastName"
            name="lastName"
            placeholder="Enter your last name"
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="companyName">
            Company Name (optional)
          </label>
          <input
            className="form-control"
            type="text"
            id="companyName"
            name="companyName"
            placeholder="Enter your company name (if any)"
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="country">
            Country/Region <span>*</span>
          </label>
          <select className="form-select" id="country" name="country" required>
            <option value="Ecuador">Ecuador</option>
          </select>
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="streetAddress">
            Street Address <span>*</span>
          </label>
          <input
            className="form-control"
            type="text"
            id="streetAddress"
            name="streetAddress"
            placeholder="Street name and house number"
            required
          />
          <input
            className="form-control"
            type="text"
            id="streetAddress2"
            name="streetAddress2"
            placeholder="Apartment, suite, etc. (optional)"
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="city">
            City <span>*</span>
          </label>
          <select className="form-select" id="city" name="city" required>
            <option value="">Select an option...</option>
          </select>
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="state">
            State/Province <span>*</span>
          </label>
          <select className="form-select" id="state" name="state" required>
            <option value="Pichincha">Pichincha</option>
          </select>
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="postalCode">
            ZIP/Postal Code <span>*</span>
          </label>
          <input
            className="form-control"
            type="text"
            id="postalCode"
            name="postalCode"
            placeholder="Enter ZIP code"
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="phone">
            Phone <span>*</span>
          </label>
          <input
            className="form-control"
            type="tel"
            id="phone"
            name="phone"
            placeholder="Enter your phone number"
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="email">
            Email Address <span>*</span>
          </label>
          <input
            className="form-control"
            type="email"
            id="email"
            name="email"
            placeholder="Enter your email address"
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label" htmlFor="orderNotes">
            Order Notes (optional)
          </label>
          <textarea
            className="form-control"
            id="orderNotes"
            name="orderNotes"
            placeholder="Notes about your order, e.g., special instructions for delivery"
            rows="3"
          ></textarea>
        </div>

        <button type="submit" className="btn btn-primary">
          Confirm Billing Details
        </button>
      </form>
    </div>
  );
}

export default BillingDetails;
