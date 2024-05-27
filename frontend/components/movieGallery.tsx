'use client'
import { useRouter } from 'next/navigation'
import { Button } from '@nextui-org/button'
import { Card, CardFooter } from '@nextui-org/card'
import { Image } from '@nextui-org/image'

import { HeartFilledIcon } from '../components/icons'

type Movie = {
  id: number
  title: string
  likes: number
  created_at: string
}

interface MovieGalleryProps {
  movies: Movie[]
}

export const MovieGallery = ({ movies }: MovieGalleryProps) => {
  const router = useRouter()

  async function likeMovie(id: number) {
    const res = await fetch(`http://localhost:3000/api/movie/${id}`, {
      method: 'PATCH',
    })
    if (!res.ok) {
      throw new Error('Failed to upvote movie');
    }

    router.refresh()
  }

  const parseMovieCards = (movies: Movie[]) =>
    movies.map((movie) => (
      <Card
        key={`movie-id-${movie.id}`}
        isFooterBlurred
        radius="md"
        className="border-none max-w-screen-60 max-h-60"
      >
        <Image
          alt="movie image default"
          className="object-cover"
          height="auto"
          src="https://nextui.org/images/hero-card.jpeg"
          width="auto"
        ></Image>
        <CardFooter className="justify-between before:bg-white/10 border-white/20 border-1 overflow-hidden py-1 absolute before:rounded-xl rounded-large bottom-1 w-[calc(100%_-_8px)] shadow-small ml-1 z-10">
          <p className="text-base text-white/100">{movie.title}</p>
          <Button
            className="text-tiny text-white bg-black/20"
            variant="flat"
            color="default"
            radius="lg"
            size="sm"
            onPress={() => likeMovie(movie.id)}
          >
            <HeartFilledIcon />
            {movie.likes}
          </Button>
        </CardFooter>
      </Card>
    ))

  const movieCards = parseMovieCards(movies)

  return <div className="grid grid-cols-3 gap-4">{movieCards}</div>
}
