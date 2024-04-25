import { useTransport } from './transport'
import { createPromiseClient } from '@connectrpc/connect'
import { AuthService } from 'src/gen/auth/v1/auth_connect'
import { BookService } from 'src/gen/book/v1/book_connect'
import { EventService } from 'src/gen/event/v1/event_connect'

export const authService = createPromiseClient(AuthService, useTransport())
export const bookService = createPromiseClient(BookService, useTransport())
export const eventService = createPromiseClient(EventService, useTransport())
