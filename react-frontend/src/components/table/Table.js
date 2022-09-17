import React from 'react'
import BootstrapTable from 'react-bootstrap-table-next'
import paginationFactory from 'react-bootstrap-table2-paginator'
import TableDetails from './TableDetails'
import { useWindowDimensions, detectMobileChange } from '../../utils/mobileResponsive'

export default function Table({data, displayRegion, highlightRegions}) {
  const { width } = useWindowDimensions()
  const [columns, setColumns] = React.useState(detectMobileChange())
  const [detailsView, setDetailsView] = React.useState(false)
  const [region, setRegion] = React.useState("")

  const defaultSorted = [
    {
      dataField: 'cases.new',
      order: 'desc'
    }
  ]

  const selectRow = {
    mode: 'radio',
    hideSelectColumn: true,
    clickToSelect: true,
    onSelect: (row, isSelect, rowIndex, e) => {
      if (region === row.country) {
        setRegion(row.country + " ")
      } else {
        setRegion(row.country)
      }
      setDetailsView(true)
    }
  }

  const rowHighlight = (row, rowIndex) => {
    const style = {}

    if (highlightRegions !== null && highlightRegions.includes(row.country)) {
      style.backgroundColor = '#c8e6c9'
    }

    return style
  }

  React.useEffect(() => {
    setColumns(detectMobileChange())
  }, [width])

  React.useEffect(() => {
    if (displayRegion !== null) {
      setRegion(displayRegion)
      setDetailsView(true)
    }
  }, [displayRegion])

  return (

      <>
        <BootstrapTable
            selectRow={ selectRow }
            bootstrap4
            keyField="country"
            data={ data }
            columns={ columns }
            rowStyle={ rowHighlight }
            defaultSorted={ defaultSorted }
            pagination={ paginationFactory() }
            striped
            hover
            condensed
        />
        { detailsView && <TableDetails region={region} /> }
      </>
    )
}