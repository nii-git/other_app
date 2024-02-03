import React from "react";
import { Link } from "react-router-dom";
import { Header } from "./Header.js";
import {
  ProgressSteps,
  NumberedStep
} from "baseui/progress-steps";
import { Button } from "baseui/button";
import styles from './About.module.css';

function AboutWithoutHeader() {
  return (
    <div>
    <header className={styles['about-header']}>
        <h1>About</h1>
    </header>

    <section className={styles['about-section']}>
        <h2>titleとは？</h2>
        <div>
          <p>ここに、サービスやチームについて簡単な紹介文を書いてください。</p>
          <p>ここに、サービスやチームについて簡単な紹介文を書いてください。</p>
        </div>

    </section>

    <section className={styles['about-section']}>
        <h2>ミッション</h2>
        <p>サービスのミッションや目標について説明してください。</p>
    </section>

    <section className={styles['about-section']}>
        <h2>チーム</h2>
        <p>チームメンバーの紹介や写真を追加してください。</p>
    </section>

    <section className={styles['about-section']}>
        <h2>お問い合わせ</h2>
        <p>何か質問やご意見があれば、お気軽にお問い合わせください。</p>
    </section>

</div>
  );
}


export const About =() => {
  return (
    <>
      <Header></Header>
      <AboutWithoutHeader></AboutWithoutHeader>
    </>
  );
}