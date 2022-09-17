import React from 'react'

const MIN_WIDTH_MOBILE = 760

export const useWindowDimensions = () => {
  const [ width, setWidth ] = React.useState(window.innerWidth)

  React.useEffect(() => {
    const resizeEvent = () => setWidth(window.innerWidth)
    window.addEventListener('resize', resizeEvent)
  })

  return { width }
}

export const isMobile = () => {
    return (window.innerWidth <= MIN_WIDTH_MOBILE)
}

export const detectMobileChange = () => {
    return (window.innerWidth <= MIN_WIDTH_MOBILE) ? columnsMobile : columnsDesktop
}

  const columnsMobile = [
    {
      dataField: 'country',
      text: 'Country',
      sort: true,
      align: 'left'
    },
    {
      dataField: 'cases.total',
      text: 'Total Cases',
      sort: true
    },
    {
      dataField: 'cases.new',
      text: 'New Cases',
      sort: true,
      sortFunc: (a, b, order, dataField) => {
        const aAsInt = (a.length > 0) ? a.substr(1) : a
        const bAsInt = (b.length > 0) ? b.substr(1) : b
        if (order === 'asc') {
          return aAsInt - bAsInt
        }
        return bAsInt - aAsInt
      }
    },
    {
      dataField: 'deaths.total',
      text: 'Total Deaths',
      sort: true
    },
    {
      dataField: 'deaths.new',
      text: 'New Deaths',
      sort: true,
      sortFunc: (a, b, order, dataField) => {
        const aAsInt = (a.length > 0) ? a.substr(1) : a
        const bAsInt = (b.length > 0) ? b.substr(1) : b
        if (order === 'asc') {
          return aAsInt - bAsInt
        }
        return bAsInt - aAsInt
      }
    }
  ]

  const columnsDesktop = [
    {
      dataField: 'decoratedCountry',
      text: 'Country',
      sort: true,
      align: 'left'
    },
    {
      dataField: 'cases.total',
      text: 'Total Cases',
      sort: true
    },
    {
      dataField: 'cases.new',
      text: 'New Cases',
      sort: true,
      sortFunc: (a, b, order, dataField) => {
        const aAsInt = (a.length > 0) ? a.substr(1) : a
        const bAsInt = (b.length > 0) ? b.substr(1) : b
        if (order === 'asc') {
          return aAsInt - bAsInt
        }
        return bAsInt - aAsInt
      }
    },
    {
      dataField: 'deaths.total',
      text: 'Total Deaths',
      sort: true
    },
    {
      dataField: 'deaths.new',
      text: 'New Deaths',
      sort: true,
      sortFunc: (a, b, order, dataField) => {
        const aAsInt = (a.length > 0) ? a.substr(1) : a
        const bAsInt = (b.length > 0) ? b.substr(1) : b
        if (order === 'asc') {
          return aAsInt - bAsInt
        }
        return bAsInt - aAsInt
      }
    },
    {
      dataField: 'cases.recovered',
      text: 'Total Recovered',
      sort: true
    },
    {
      dataField: 'cases.active',
      text: 'Active Cases',
      sort: true
    },
    {
      dataField: 'cases.critical',
      text: 'Serious, Critical',
      sort: true
    },
    {
      dataField: 'tests.total',
      text: 'Total Tests',
      sort: true
    }
  ]