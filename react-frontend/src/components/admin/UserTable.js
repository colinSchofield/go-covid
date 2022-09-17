import React from 'react'
import BootstrapTable from 'react-bootstrap-table-next'
import SignUp from '../signup/SignUp'

export default function UserTable({data, clearTableData}) {
  const [userId, setUserId] = React.useState(null)

  const userColumns = [
        {
          dataField: 'id',
          text: 'Id',
          hidden: true
        },
        {
          dataField: 'name',
          text: 'Name',
        },
        {
          dataField: 'age',
          text: 'Age',
        },
        {
          dataField: 'gender',
          text: 'Gender',
        },
        {
          dataField: 'contact',
          text: 'Contact',
        },
        {
          dataField: 'regionList',
          text: 'Regions',
        }
      ]

  const defaultSorted = [
    {
      dataField: 'age',
      order: 'desc'
    }
  ]

  const selectRow = {
    mode: 'radio',
    hideSelectColumn: true,
    clickToSelect: true,
    onSelect: (row, isSelect, rowIndex, e) => {
      setUserId(row.id)
    }
  }

  return (
      <>
        { userId && <SignUp adminId={userId} returnToAdminTable={clearTableData} /> }

        { !userId &&
          <>
            <br/>
            <h4>Signed-Up Users</h4>
            <i>(Usually, <b>Admin access</b> requires some form of authentication, but just this once..)</i>
            <br/><br/>
            <p class="font-weight-bold blue-text">You may view, edit or delete users from this table.</p>
            <BootstrapTable
                selectRow={ selectRow }
                bootstrap4
                keyField="id"
                data={ data }
                columns={ userColumns }
                defaultSorted={ defaultSorted }
                striped
                condensed
                hover />
          </>
        }
      </>
    )
}