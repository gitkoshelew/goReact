import { Alert, AlertTitle, Snackbar } from '@mui/material'
import { useState } from 'react'

export type LoginErrorMsgPropsType = {
  ErrorMsg: string
}

export const LoginErrorMsg = ({ ErrorMsg }: LoginErrorMsgPropsType) => {
  const [openSnack, setOpenShack] = useState<boolean>(true)
  const handleClose = () => {
    setOpenShack(false)
  }
  return (
    <Snackbar
      open={openSnack}
      onClose={handleClose}
      autoHideDuration={4000}
      anchorOrigin={{ vertical: 'top', horizontal: 'left' }}
    >
      <Alert onClose={handleClose} variant="filled" severity="error" sx={{ width: '100%' }}>
        <AlertTitle>Error</AlertTitle>
        <strong>{ErrorMsg}</strong>
      </Alert>
    </Snackbar>
  )
}
