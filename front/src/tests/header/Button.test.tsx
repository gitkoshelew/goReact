import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { BasketButton } from '../../main/ui/components/Button/BasketButton/BasketBtn'
import { LoginButton } from '../../main/ui/components/Button/LoginBtn/LoginBtn'

describe('Button', () => {
  it('if transferred /login/ type should contain login div', () => {
    render(
      <BrowserRouter>
        <LoginButton />
      </BrowserRouter>
    )
    expect(screen.getByText(/login/i)).toBeInTheDocument()
  })
  it('if transferred /buy/ type should contain img elem with alt=buyContainer', () => {
    render(
      <BrowserRouter>
        <BasketButton />
      </BrowserRouter>
    )

    expect(screen.getByRole('img')).toBeInTheDocument()
    expect(screen.getByAltText('buyContainer')).toBeInTheDocument()
  })
})
