import React from 'react'
import Breadcrumb from "@/components/Breadcrumb";

function CategoryLayout({children}) {
  return (
    <>
    <Breadcrumb />
    {children}
    </>
  )
}

export default CategoryLayout
