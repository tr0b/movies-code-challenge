import { NextResponse } from "next/server"

export async function PATCH(
  request: Request,
  { params }: { params: { id: string } },
) {
  const id = params.id
  console.log(id)
  const res = await fetch(`http://api:8080/movie/${id}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    method: 'PATCH',
    body: JSON.stringify({
      action: 'UPVOTE',
    }),
  })
  const data = await res.json()

  return NextResponse.json({ data })
}
