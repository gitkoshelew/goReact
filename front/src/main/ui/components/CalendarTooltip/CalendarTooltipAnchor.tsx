import moment from 'moment'
import React from 'react'
import s from './CalendarTooltipAnchor.module.scss'

type PropsType = {
  date: Date
  showTooltip: (isVisible: boolean) => void
  handleTooltipDate: (date: Date | null) => void
}

export const CalendarTooltipAnchor = React.memo(({ date, showTooltip, handleTooltipDate }: PropsType) => {
  const handleShowTooltip = () => {
    const prevDay = moment().add(-1, 'days')
    const selectedDay = moment(date)
    if (selectedDay >= prevDay) {
      showTooltip(true)
      handleTooltipDate(date)
    }
  }
  const handleHideTooltip = () => {
    showTooltip(false)
    handleTooltipDate(null)
  }

  return <div className={s.tooltipAnchor} onMouseEnter={handleShowTooltip} onMouseLeave={handleHideTooltip} />
})
