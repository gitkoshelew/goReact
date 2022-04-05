import { render } from '@testing-library/react'
import { Footer } from '../../main/ui/pages/footer/footer'
import { BrowserRouter } from 'react-router-dom'

test('footer should contain useFullLinks && stayInTouch block', () => {
  const { getByText } = render(
    <BrowserRouter>
      <Footer />
    </BrowserRouter>
  )
  //is useFullLinks block exist
  const useFullLinksBlockCheck = getByText(/USEFUL LINKS/i)
  const useFullLinksElem = getByText(/rooms/i)
  expect(useFullLinksBlockCheck).toBeInTheDocument()
  expect(useFullLinksElem).toBeInTheDocument()
  //is stayInTouchBlock exist
  const stayInTouchBlockCheck = getByText(/STAY IN TOUCH/i)
  expect(stayInTouchBlockCheck).toBeInTheDocument()
})
