import { useTransport } from './transport'
import { createClient } from '@connectrpc/connect'
import { AuthService } from 'src/gen/auth/v1/auth_pb'
import { BookService } from 'src/gen/book/v1/book_pb'
import { EventService } from 'src/gen/event/v1/event_pb'

export const authService = createClient(AuthService, useTransport())
export const bookService = createClient(BookService, useTransport())
export const eventService = createClient(EventService, useTransport())
