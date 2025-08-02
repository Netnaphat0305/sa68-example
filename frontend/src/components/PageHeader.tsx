// src/components/PageHeader.tsx
import React from 'react';
import styles from './PageHeader.module.css'; // ไฟล์ CSS module สำหรับตกแต่ง

type PageHeaderProps = {
  title: string;
};

const PageHeader: React.FC<PageHeaderProps> = ({ title }) => {
  return <h1 className={styles.pageHeader}>{title}</h1>;
};

export default PageHeader;
