import { Button } from '@nextui-org/button'
import { Input } from '@nextui-org/react'
import { redirect } from 'next/navigation'

export const NewMovieForm = () => {
  async function createMovie(formData: FormData) {
    'use server'
    const title: FormDataEntryValue | null = formData.get('title')

    if (!title) {
      return;
    }

    const res = await fetch(`http://frontend:3000/api/movie/`, {
      method: 'POST',
      body: JSON.stringify({ title: title.toString() }),
    })

    if (!res.ok) {
      throw new Error('Failed to submit form')
    }

    redirect('/')
  }

  return (
    <>
      <br />
      <form action={createMovie}>
        <div className="flex w-full flex-wrap md:flex-nowrap gap-4">
          <Input
            type="text"
            name="title"
            label="Movie Title"
            aria-label="Movie Title"
            placeholder="Enter your Movie Title"
          />
        </div>
        <br />
        <Button color="default" type="submit">
          Add
        </Button>
      </form>
    </>
  )
}
