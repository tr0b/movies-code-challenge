import { NextResponse } from "next/server"

export async function GET(request: Request) {
  const res = await fetch('http://api:8080/movie', {
    headers: {
      'Content-Type': 'application/json',
    },
  })
  const movies = await res.json()

  return NextResponse.json({ movies })
}

export async function POST(request: Request) {
  const body = await request.json()
  console.log('body of req', body)
  const res = await fetch('http://api:8080/movie', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ title: body?.title }),
  })
  const data = await res.json()
  console.log(data)

  return NextResponse.json({ data })
}
