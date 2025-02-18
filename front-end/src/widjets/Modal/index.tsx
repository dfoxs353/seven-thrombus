import { useEffect, useRef, ReactNode } from 'react';
import { Typography } from '@shared/ui';
import { ReactSVG } from 'react-svg';
import back from '@/assets/images/icons/back.svg'

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  title: string;
  children: ReactNode;
}

export const Modal = ({ isOpen, onClose, title, children }: ModalProps) => {
  const dialogRef = useRef<HTMLDialogElement>(null);

  useEffect(() => {
    const dialog = dialogRef.current;
    if (!dialog) return;

    if (isOpen) {
      dialog.showModal();
    } else {
      dialog.close();
    }
  }, [isOpen]);

  const handleBackdropClick = (e: React.MouseEvent<HTMLDialogElement>) => {
    if (e.target === dialogRef.current) {
      onClose();
    }
  };

  return (
    <dialog
      ref={dialogRef}
      onClose={onClose}
      onClick={handleBackdropClick}
      className="modal"
    >
      <div className="p-4">
        <div className="flex items-center gap-4 mb-5">
          <ReactSVG onClick={onClose} src={back} className='p-1.5 cursor-pointer'/>
          <Typography as='h2' className='heading'>{title}</Typography>
        </div>
        <div className="">{children}</div>
      </div>
    </dialog>
  );
};