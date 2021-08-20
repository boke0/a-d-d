import styles from '../../styles/atm/Btn.module.scss';

export default function Btn({ children, outline = false, ...props }) {
  return (
    <button className={`${styles.button} ${outline ? styles.outline : ''}`} {...props}>
      <span>{children}</span>
    </button>
  )
}
