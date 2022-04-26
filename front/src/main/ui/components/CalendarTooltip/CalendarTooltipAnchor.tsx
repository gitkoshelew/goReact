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
    const nextMaxDate = moment().add(+1, 'month')
    const prevDay = moment().add(-1, 'days')
    const selectedDay = moment(date)
    if (selectedDay >= prevDay && selectedDay <= nextMaxDate) {
      showTooltip(true)
      handleTooltipDate(date)
    }
    console.log(date)
    if (date >= new Date()) {
      console.log(true)
      // Дата на которую наведен курсор - сегодняшняя дата -
    }
  }

  //в юзстейт записать ID
  //посчитать на каком я дне. 26 + ID4 = 30.
  const handleHideTooltip = () => {
    showTooltip(false)
    handleTooltipDate(null)
  }

  return <div className={s.tooltipAnchor} onMouseEnter={handleShowTooltip} onMouseLeave={handleHideTooltip} />
})
