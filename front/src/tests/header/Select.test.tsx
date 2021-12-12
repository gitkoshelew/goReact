import { render, screen } from '@testing-library/react'
import { SelectUI } from '../../main/ui/components/Select/Select'

describe('SelectUI', () => {
  it('Select should contain 3 option', () => {
    render(<SelectUI />)

    expect(screen.queryAllByRole('option')).toHaveLength(3)
  })
})
