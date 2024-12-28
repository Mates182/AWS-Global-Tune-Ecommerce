import React from 'react'
import Link from "next/link";

function NotFound() {
  return (
    <section>
        <h1 className='text-center'>404 Hz: Tune not found</h1>
      <h3>You seem to be lost...</h3>
      <Link href='/' className='btn btn-light'>Return to Home</Link>
    </section>
  )
}

export default NotFound
