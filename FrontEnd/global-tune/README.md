This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

# Getting Started


## Local Development Setup

To run the project locally, ensure you have a `.env.local` file in the root of the frontend project (`/globaltune`). This file will be used to manage the environment variables needed for the application. 

### Example `.env.local` File

Create a `.env.local` file and define the following variable:

```plaintext
GRAPHQL_PRODUCTS_ENDPOINT="http://your-graphql-endpoint/graphql/"
```

Replace `http://your-graphql-endpoint/graphql/` with the actual URL of your GraphQL endpoint.

### Steps to Run Locally

To run the development server:

```bash
npm install
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

