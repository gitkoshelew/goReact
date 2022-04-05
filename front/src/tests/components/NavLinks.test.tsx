import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { NavLinks } from '../../main/ui/components/NavLinks/NavLinks'
import { NavLinksBurger } from '../../main/ui/components/NavLinks/NavLinksBurger/NavLinksBurger'
import userEvent from '@testing-library/user-event'

describe('NavLinks', () => {
  it('NavLinks should contain all transferred array of string in link tag', () => {
    const testArr = ['test0', 'test1', 'test2', 'test3']
    render(
      <BrowserRouter>
        <NavLinks navNames={testArr} />
      </BrowserRouter>
    )

    expect(screen.queryAllByRole(/link/i)).toHaveLength(testArr.length)
  })
  it('NavLinksBurger should contain all transferred array of string in link tag', () => {
    const testArr = ['test0', 'test1', 'test2', 'test3', 'test4']
    render(
      <BrowserRouter>
        <NavLinksBurger navNames={testArr} />
      </BrowserRouter>
    )
    expect(screen.queryAllByRole(/link/i)).toHaveLength(testArr.length)
  })
  it('NavLink (in burger instance) should become active(change className) when user has click on it', () => {
    const testArr = ['test0', 'test1', 'test2', 'test3', 'test4']
    render(
      <BrowserRouter>
        <NavLinksBurger navNames={testArr} />
      </BrowserRouter>
    )

    const oneLink = screen.getAllByRole(/link/i)[0]
    expect(oneLink.classList.contains('oneLinkActive')).toBe(false)

    userEvent.click(oneLink)

    expect(oneLink.classList.contains('oneLinkActive')).toBe(true)
  })
  it('NavLink (in native instance) should become active(change className) when user has click on it', () => {
    const arr = ['test0', 'test1', 'test2', 'test3', 'test4']

    render(
      <BrowserRouter>
        <div>
          <NavLinks navNames={arr} />
        </div>
      </BrowserRouter>
    )

    const oneLink = screen.getAllByRole(/link/i)[0]
    expect(oneLink.classList.contains('oneLinkActive')).toBe(true)

    const secondLink = screen.getAllByRole(/link/i)[1]
    userEvent.click(secondLink)

    expect(oneLink.classList.contains('oneLinkActive')).toBe(false)
    expect(secondLink.classList.contains('oneLinkActive')).toBe(true)
  })
})
