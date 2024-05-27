import { title } from '@/components/primitives'
import { MovieGallery } from '../components/movieGallery'

async function getMovies() {
  const res = await fetch('http://frontend:3000/api/movie', {
    cache: 'no-store',
  })

  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }

  const movies = await res.json().then((data) => data.movies)
  console.log(movies)
  return movies
}
export default async function Home() {
  const movies = await getMovies()
  return (
    <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
      <div className="inline-block max-w-lg text-center justify-center">
        <h1 className={title()}>Movie Store</h1>
        <MovieGallery movies={movies} />
      </div>
    </section>
  )
}
