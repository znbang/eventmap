import { useTransport } from './transport'
import { createPromiseClient } from '@bufbuild/connect-web'
import { AuthService } from 'src/gen/auth/v1/auth_connectweb'
import { BookService } from 'src/gen/book/v1/book_connectweb'
import { EventService } from 'src/gen/event/v1/event_connectweb'

export const authService = createPromiseClient(AuthService, useTransport())
export const bookService = createPromiseClient(BookService, useTransport())
export const eventService = createPromiseClient(EventService, useTransport())
