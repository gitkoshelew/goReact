import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { Button } from '../../main/ui/components/Button/Button'

describe('Button', () => {
  it('if transferred /login/ type should contain login div', () => {
    render(
      <BrowserRouter>
        <Button type={'login'} />
      </BrowserRouter>
    )
    screen.debug()
    expect(screen.getByText(/login/i)).toBeInTheDocument()
  })
  it('if transferred /buy/ type should contain img elem with alt=buyContainer', () => {
    render(
      <BrowserRouter>
        <Button type={'buy'} />
      </BrowserRouter>
    )
    screen.debug()
    expect(screen.getByRole('img')).toBeInTheDocument()
    expect(screen.getByAltText('buyContainer')).toBeInTheDocument()
  })
})
