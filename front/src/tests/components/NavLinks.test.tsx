import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { NavLinks } from '../../main/ui/components/NavLinks/NavLinks'
import { NavLinksBurger } from '../../main/ui/components/NavLinks/NavLinksBurger/NavLinksBurger'

describe('NavLinks', () => {
  it('NavLinks should contain all transferred array of string in link tag', () => {
    const testArr = ['test0', 'test1', 'test2', 'test3']
    render(
      <BrowserRouter>
        <NavLinks navNames={testArr} />
      </BrowserRouter>
    )
    screen.debug()
    expect(screen.queryAllByRole(/link/i)).toHaveLength(testArr.length)
  })
  it('NavLinksBurger should contain all transferred array of string in link tag', () => {
    const testArr = ['test0', 'test1', 'test2', 'test3', 'test4']
    render(
      <BrowserRouter>
        <NavLinksBurger navNames={testArr} />
      </BrowserRouter>
    )
    screen.debug()
    expect(screen.queryAllByRole(/link/i)).toHaveLength(testArr.length)
  })
})
