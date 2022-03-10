import { useTransport } from './transport'
import { createPromiseClient } from '@bufbuild/connect-web'
import { AuthService } from '../../gen/auth/v1/auth_connectweb'
import { BookService } from '../../gen/book/v1/book_connectweb'
import { EventService } from '../../gen/event/v1/event_connectweb'

export const authService = createPromiseClient(AuthService, useTransport())
export const bookService = createPromiseClient(BookService, useTransport())
export const eventService = createPromiseClient(EventService, useTransport())
