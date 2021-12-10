import { Carousel } from '../../main/ui/components/Carousel/Carousel'
import { render, screen } from '@testing-library/react'

describe('carousel', () => {
  it('carousel should contain transferred children element', () => {
    const childrenElement = (
      <div data-testid="test">
        <div>test</div>
      </div>
    )

    render(
      <Carousel startSlideToShow={1} startSlideScroll={1} type={'smallWidth'} isWithArrows={true}>
        {childrenElement}
      </Carousel>
    )
    screen.debug()
    expect(screen.getByTestId('test')).toBeInTheDocument()
  })
})
