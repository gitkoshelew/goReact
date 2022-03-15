import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { BurgerMenu } from '../../main/ui/pages/navBar/burgerMenu/BurgerMenu'
import { useState } from 'react'
import userEvent from '@testing-library/user-event'

describe('BurgerMenu', () => {
  it('BurgerMenu should show only 8 links whithout burgerMenu', () => {
    render(
      <BrowserRouter>
        <BurgerMenu setIsBurgerCollapse={() => {}} isBurgerCollapse={true} />
      </BrowserRouter>
    )
    expect(screen.getAllByRole('link')).toHaveLength(9)
  })

  it('BurgerMenu doesnt show any links with negative isBurgerCollapse', () => {
    render(
      <BrowserRouter>
        <BurgerMenu setIsBurgerCollapse={() => {}} isBurgerCollapse={false} />
      </BrowserRouter>
    )

    const link = screen.queryAllByRole(/link/i)
    expect(link).toHaveLength(0)
  })
  it('BurgerMenu should open when user click on burger icon', () => {
    const Test = () => {
      const [isBurgerCollapse, setIsBurgerCollapse] = useState(false)
      return <BurgerMenu setIsBurgerCollapse={setIsBurgerCollapse} isBurgerCollapse={isBurgerCollapse} />
    }
    render(
      <BrowserRouter>
        <Test />
      </BrowserRouter>
    )
    const burger = screen.getByTestId('test')

    const beforeClickLink = screen.queryAllByRole(/link/i)
    expect(beforeClickLink).toHaveLength(0)

    userEvent.click(burger)

    const afterClickLink = screen.queryAllByRole(/link/i)
    expect(afterClickLink).toHaveLength(9)
  })
})
