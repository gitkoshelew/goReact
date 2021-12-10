import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { BurgerMenu } from '../../main/ui/pages/navBar/burgerMenu/BurgerMenu'

describe('BurgerMenu', () => {
  it('BurgerMenu should show only 8 links whithout burgerMenu', () => {
    render(
      <BrowserRouter>
        <BurgerMenu setIsBurgerCollapse={() => {}} isBurgerCollapse={true} />
      </BrowserRouter>
    )
    screen.debug()
    expect(screen.getAllByRole('link')).toHaveLength(8)
  })

  it('BurgerMenu doesnt show any links with negative isBurgerCollapse', () => {
    render(
      <BrowserRouter>
        <BurgerMenu setIsBurgerCollapse={() => {}} isBurgerCollapse={false} />
      </BrowserRouter>
    )
    screen.debug()
    const link = screen.queryAllByRole(/link/i)
    expect(link).toHaveLength(0)
  })
})
