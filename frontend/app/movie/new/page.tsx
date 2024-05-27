import { NewMovieForm } from '../../../components/NewMovieForm'

// `app/dashboard/page.tsx` is the UI for the `/dashboard` URL
export default function Page() {
  return (
    <>
      <h1>New Movie</h1>
      <NewMovieForm />
    </>
  )
}
