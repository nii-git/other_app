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
        <h2>FrequENとは？</h2>
        <div>
          <p>英語ニュースサイトに登場する英単語の出現頻度を分析するサイトです。</p>
          <p>サービス名はFrequency(頻度)+ English を組み合わせたものです。</p>
          <p>タイトルはそのうち変わるかもしれません。</p>
        </div>

    </section>

    <section className={styles['about-section']}>
        <h2>どんな時に使う？</h2>
        <div>
          <h3>出現頻度の高い英単語を知りたい時</h3>
          <p>SAMPLEは、単語の出現頻度を知ることができます。</p>
          <h3>トレンドの英単語を知りたい時</h3>
          <p>英字ニュースを見るにあたって、トレンドである単語を抑えることは大事です。</p>
          <p>急に上位にあがってきた英単語をチェックするのに活用してください。</p>
          <p>そのうち急上昇単語機能を実装します。たぶん。</p>
          <h3>英語ニュースサイトを探したい時</h3>
          <p>「Websites」一覧から、あなたに合うニュースサイトを探してみてください。</p>
          <h3>英語学習の単語帳にも</h3>
          <p>滅多に出てこないレア単語を覚えるより、頻出単語から覚えた方が効率的です。</p>
          <p>英語学習者にも使っていただくため、そのうち日本語の意味を表示できるようにします。たぶん。</p>
        </div>

    </section>

    <section className={styles['about-section']}>
        <h2>推奨環境</h2>
        <p>PC版 Google Chromeを推奨しています。</p>
        <p>それ以外の環境は動作確認を行なっていません。</p>
        <p style={{color:"red"}}>モバイル環境は非推奨です。</p>
    </section>

    <section className={styles['about-section']}>
        <h2>お問い合わせ</h2>
        <p>X(Twitter): <a href="https://twitter.com/nii_tech">@nii_tech</a> までお願いします。</p>
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