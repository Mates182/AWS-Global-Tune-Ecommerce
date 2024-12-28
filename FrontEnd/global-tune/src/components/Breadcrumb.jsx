"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

// Function to transform text to "Capitalized Case"
function capitalizeFirstLetter(str) {
  return str
    .toLowerCase()
    .split("-")
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(" ");
}

function Breadcrumb() {
  const pathname = usePathname();
  // TODO: split and exclude a given path
  // Split the pathname into parts, excluding "category"
  const pathParts = pathname
    .split("/")
    .filter((part) => part && part !== "category" && part !== "product");

  // If no valid parts are found, do not render the breadcrumb
  if (pathParts.length === 0) {
    return null;
  }
  // TODO: use the breadcrumb dinamicaly with different directions, now only uses Category
  return (
    <nav aria-label="breadcrumb">
      <ol className="breadcrumb">
        {/* Home link */}
        <li className="breadcrumb-item">
          <Link href="/">Home</Link>
        </li>
        {/* Dynamic breadcrumb links */}
        {pathParts.map((part, index) => {
          const href = "/category/" + pathParts.slice(0, index + 1).join("/");
          const isLast = index === pathParts.length - 1;

          return (
            <li
              key={href}
              className={`breadcrumb-item${isLast ? " active" : ""}`}
              aria-current={isLast ? "page" : undefined}
            >
              {isLast ? (
                capitalizeFirstLetter(part) // Display capitalized text for the last breadcrumb
              ) : (
                <Link href={href}>{capitalizeFirstLetter(part)}</Link> // Link for intermediate breadcrumbs
              )}
            </li>
          );
        })}
      </ol>
    </nav>
  );
}

export default Breadcrumb;
